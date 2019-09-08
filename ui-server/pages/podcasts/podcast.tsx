import EpisodeList from 'components/episode_list'
import LoadingPage from 'components/loading_page'
import PodcastDetails from 'components/podcast_details'
import React, { Component } from 'react'
import { RequestState } from 'reducers/requests/utils'
import { PageContext } from 'types/utilities'

export interface StateToProps {
  reqState: RequestState
}

export interface DispatchToProps {
  loadPodcast: (id: string) => void
}

export interface OwnProps {
  podcastId: string
  preventInitialLoad: boolean
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

export default class PodcastPage extends Component<Props> {
  static async getInitialProps({ query }: PageContext) {
    const id = query['id'] as string
    return { podcastId: id, preventInitalLoad: false }
  }

  componentDidUpdate(prevProps: Props) {
    if (this.props.podcastId != prevProps.podcastId) {
      this.props.loadPodcast(this.props.podcastId)
    }
  }

  componentDidMount() {
    if (!this.props.preventInitialLoad) {
      this.props.loadPodcast(this.props.podcastId)
    }
  }

  render() {
    const { reqState, podcastId } = this.props

    if (reqState.status == 'STARTED' || reqState.status == 'NOT_STARTED') {
      return <LoadingPage />
    }

    if (reqState.status == 'SUCCESS') {
      return (
        <>
          <PodcastDetails podcastId={podcastId} />
          <EpisodeList podcastId={podcastId} />
        </>
      )
    }

    return <></>
  }
}
