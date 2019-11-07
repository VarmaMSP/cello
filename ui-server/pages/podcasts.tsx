import { getPodcast } from 'actions/podcast'
import PodcastView from 'components/podcast_view/podcast_view'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface OwnProps {
  podcastId: string
  context: string
  scrollY: number
}

export default class PodcastsPage extends Component<OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { query, store } = ctx
    await bindActionCreators(getPodcast, store.dispatch)(query[
      'podcastId'
    ] as string)
  }

  componentDidMount() {
    gtag.pageview(`/podcasts/${this.props.podcastId}/${this.props.context}`)

    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    return <PodcastView podcastId={this.props.podcastId} />
  }
}
