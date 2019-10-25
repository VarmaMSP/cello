import { getUserFeed } from 'actions/user'
import ButtonSignin from 'components/button_signin'
import { iconMap } from 'components/icon'
import ListFeed from 'components/list_feed'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { AppState } from 'store'
import { AppActions, SET_CURRENT_URL_PATH } from 'types/actions'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface StateToProps {
  isUserSignedIn: boolean
}

interface DispatchToProps {
  loadFeed: () => void
}

interface OwnProps {
  scrollY: number
}

class FeedPage extends React.Component<
  StateToProps & DispatchToProps & OwnProps
> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store } = ctx

    if (getIsUserSignedIn(store.getState())) {
      await bindActionCreators(getUserFeed, store.dispatch)()
    }
    store.dispatch({ type: SET_CURRENT_URL_PATH, urlPath: '/feed' })
  }

  componentDidMount() {
    gtag.pageview('/feed')

    this.props.loadFeed()
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isUserSignedIn } = this.props

    if (!isUserSignedIn) {
      const FeedIcon = iconMap['feed']
      return (
        <>
          <div className="mx-auto mt-24">
            <FeedIcon className="w-12 h-12 mx-auto fill-current text-gray-700" />
            <h1 className="text-center text-xl text-gray-700 my-6">
              {'Get Feed from your subscriptions'}
            </h1>
            <div className="w-32 mx-auto">
              <ButtonSignin />
            </div>
          </div>
        </>
      )
    }

    return (
      <>
        <NextSeo
          noindex
          title="Feed - Phenopod"
          description="Feed"
          canonical="https://phenopod.com/feed"
        />
        <ListFeed />
      </>
    )
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    isUserSignedIn: getIsUserSignedIn(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadFeed: bindActionCreators(getUserFeed, dispatch),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(FeedPage)
