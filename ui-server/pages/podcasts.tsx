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
import * as gtag from 'utils/gtag'

interface StateToProps {
  reqState: RequestState
}

interface OwnProps {
  podcastId: string
  scrollY: number
}

class PodcastsPage extends Component<StateToProps & OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { query, store } = ctx
    await bindActionCreators(getPodcast, store.dispatch)(query[
      'podcastId'
    ] as string)
  }

  componentDidMount() {
    gtag.pageview(`/podcasts/${this.props.podcastId}`)

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
          <div className="lg:w-4/6 w-full">
            <PodcastDetails podcastId={podcastId} />
          </div>
          <div className="flex mt-8 mb-4 ">
            <div className="bg-green-200 mr-4 px-3 py-1 text-sm rounded-full">
              Episodes
            </div>
            <div className="bg-green-200 px-3 py-1 text-sm rounded-full">
              About
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
