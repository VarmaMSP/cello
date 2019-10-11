import { getPodcast } from 'actions/podcast'
import ListEpisodes from 'components/list_episodes'
import LoadingPage from 'components/loading_page'
import PodcastDetails from 'components/podcast_details'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { AppState } from 'store'
import { PageContext } from 'types/utilities'
import { logPageView } from 'utils/analytics'
export interface StateToProps {
  reqState: RequestState
}

export interface OwnProps {
  podcastId: string
  scrollY: number
}

class PodcastsPage extends Component<StateToProps & OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { query, store, isServer } = ctx
    const loadPodcast = bindActionCreators(getPodcast, store.dispatch)(query[
      'podcastId'
    ] as string)

    if (isServer) {
      await loadPodcast
    }
  }

  componentDidMount() {
    logPageView()

    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState, podcastId } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }

    if (reqState.status == 'SUCCESS') {
      return (
        <div>
          <PodcastDetails podcastId={podcastId} />
          <div className="flex mt-8 mb-2">
            <div className="bg-green-200 px-3 py-1 md:text-base text-sm rounded-full">
              Episodes
            </div>
          </div>
          <div className="lg:w-4/6 w-full">
            <ListEpisodes podcastId={podcastId} />
          </div>
        </div>
      )
    }

    return <></>
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.podcast.getPodcast,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  PodcastsPage,
)
