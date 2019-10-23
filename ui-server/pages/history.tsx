import { getEpisodePlaybackHistory } from 'actions/episode'
import ListHistory from 'components/list_history'
import LoadingPage from 'components/loading_page'
import { NextSeo } from 'next-seo'
import React from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions, SET_CURRENT_URL_PATH } from 'types/actions'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface StateToProps {
  reqState: RequestState
}

interface DispatchToProps {
  loadHistory: () => void
}

interface OwnProps {
  scrollY: number
}

class FeedPage extends React.Component<
  StateToProps & DispatchToProps & OwnProps
> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    ctx.store.dispatch({ type: SET_CURRENT_URL_PATH, urlPath: '/history' })
  }

  componentDidMount() {
    gtag.pageview('/history')

    this.props.loadHistory()
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }
    if (reqState.status == 'SUCCESS') {
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
    return <></>
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.episode.getPlaybackHistory,
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
