import { getPodcastPageData } from 'actions/podcast'
import PodcastView from 'components/podcast_view/podcast_view'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

interface OwnProps {
  podcastUrlParam: string
  activeTab: string
  scrollY: number
}

export default class PodcastsPage extends Component<OwnProps> {
  static async getInitialProps({ query, store }: PageContext): Promise<void> {
    if (!!query['skipLoad']) {
      return
    }

    await bindActionCreators(
      getPodcastPageData,
      store.dispatch,
    )(query['podcastUrlParam'] as string)
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    return (
      <PodcastView
        podcastUrlParam={this.props.podcastUrlParam}
        activeTab={this.props.activeTab}
      />
    )
  }
}
