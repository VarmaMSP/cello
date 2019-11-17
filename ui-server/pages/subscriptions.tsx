import { getUserFeed } from 'actions/user'
import ButtonSignin from 'components/button_signin'
import { iconMap } from 'components/icon'
import SubscriptionsView from 'components/subscriptions_view/subscriptions_view'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { AppState } from 'store'
import { AppActions, SET_CURRENT_URL_PATH } from 'types/actions'
import { PageContext } from 'types/utilities'
import { now } from 'utils/format'
import * as gtag from 'utils/gtag'

interface StateToProps {
  isUserSignedIn: boolean
}

interface DispatchToProps {
  loadFeed: (publishedBefore: string) => void
}

interface OwnProps {
  scrollY: number
}

type Props = StateToProps & DispatchToProps & OwnProps

class SubscriptionsPage extends React.Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store } = ctx

    if (getIsUserSignedIn(store.getState())) {
      await bindActionCreators(getUserFeed, store.dispatch)(now())
    }
    store.dispatch({ type: SET_CURRENT_URL_PATH, urlPath: '/subscriptions' })
  }

  componentDidUpdate(prevProps: Props) {
    const { isUserSignedIn } = this.props
    if (isUserSignedIn && !prevProps.isUserSignedIn) {
      this.props.loadFeed(now())
    }
  }

  componentDidMount() {
    gtag.pageview('/subscriptions')
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isUserSignedIn } = this.props

    if (!isUserSignedIn) {
      const SubscribeIcon = iconMap['heart']
      return (
        <>
          <div className="mx-auto mt-24">
            <SubscribeIcon className="w-12 h-12 mx-auto fill-current text-gray-700" />
            <h1 className="text-center text-xl text-gray-700 my-6">
              {'Sign in to subscribe to your favourite podcasts'}
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
          title="Subscriptions - Phenopod"
          description="Subscriptions"
          canonical="https://phenopod.com/subscripitions"
        />
        <SubscriptionsView />
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
    loadFeed: (publishedBefore: string) =>
      bindActionCreators(getUserFeed, dispatch)(publishedBefore),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(SubscriptionsPage)
