import React, { Component } from 'react'
import { PageContext } from 'types/next'
import { Podcast, Episode } from 'types/app'
import PodcastDetails from '../../components/podcast_details'
import EpisodeList from '../../components/episode_list'

export interface StateToProps {
  podcast: Podcast
  episodes: Episode[]
}

export interface DispatchToProps {
  getPodcast: (id: string) => void
}

export interface OwnProps {
  id: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

export default class PodcastPage extends Component<Props> {
  static async getInitialProps({ query }: PageContext) {
    const id = query['id'] as string
    return { id }
  }

  componentDidUpdate(prevProps: Props) {
    if (this.props.id != prevProps.id) {
      this.props.getPodcast(this.props.id)
    }
  }

  componentDidMount() {
    this.props.getPodcast(this.props.id)
  }

  render() {
    const { podcast, episodes } = this.props
    if (!podcast) {
      return (
        <>
          <div>Im Mr MeeSeeks Look at me</div>
          <div>Im Loading....</div>
        </>
      )
    }

    return (
      <>
        <PodcastDetails
          title={podcast.title}
          author={podcast.author}
          description={podcast.description}
        />
        <EpisodeList episodes={episodes} />
      </>
    )
  }
}
