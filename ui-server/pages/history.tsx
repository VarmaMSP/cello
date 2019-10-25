import { getEpisodePlaybackHistory } from 'actions/episode'
import ButtonSignin from 'components/button_signin'
import { iconMap } from 'components/icon'
import ListHistory from 'components/list_history'
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
  loadHistory: () => void
}

interface OwnProps {
  scrollY: number
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

class FeedPage extends React.Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store } = ctx

    if (getIsUserSignedIn(store.getState())) {
      await bindActionCreators(getEpisodePlaybackHistory, store.dispatch)()
    }
    store.dispatch({ type: SET_CURRENT_URL_PATH, urlPath: '/history' })
  }

  componentDidUpdate(prevProps: Props) {
    const { isUserSignedIn } = this.props
    if (isUserSignedIn && !prevProps.isUserSignedIn) {
      this.props.loadHistory()
    }
  }

  componentDidMount() {
    gtag.pageview('/history')
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isUserSignedIn } = this.props

    if (!isUserSignedIn) {
      const HistoryIcon = iconMap['history']
      return (
        <>
          <div className="mx-auto mt-24">
            <HistoryIcon className="w-12 h-12 mx-auto fill-current text-gray-700" />
            <h1 className="text-center text-xl text-gray-700 my-6">
              {'Sign in to keep track of what you listen'}
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
          title="History - Phenopod"
          description="History"
          canonical="https://phenopod.com/feed"
        />
        <ListHistory />
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
    loadHistory: bindActionCreators(getEpisodePlaybackHistory, dispatch),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(FeedPage)
