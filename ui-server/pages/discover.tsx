import { getPodcastsInList } from 'actions/discover'
import DiscoverView from 'components/discover_view/discover_view'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface OwnProps {
  listId: string
  scrollY: number
}

export default class DiscoverPage extends Component<OwnProps> {
  static async getInitialProps({ query, store }: PageContext): Promise<void> {
    await bindActionCreators(
      getPodcastsInList,
      store.dispatch,
    )(query['listId'] as string)
  }

  componentDidMount() {
    gtag.pageview(`/discover/${this.props.listId}`)
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    return (
      <div>
        <DiscoverView listId={this.props.listId} />
      </div>
    )
  }
}
