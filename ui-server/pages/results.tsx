import { searchPodcasts } from 'actions/search'
import ListSearchResults from 'components/list_search_results'
import { NextSeo } from 'next-seo'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface OwnProps {
  query: string
  scrollY: number
}

export default class ResultsPage extends Component<OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { query, store } = ctx
    await bindActionCreators(
      searchPodcasts,
      store.dispatch,
    )(query['query'] as string)
  }

  componentDidMount() {
    gtag.search(this.props.query)
    gtag.pageview(`/results?query=${this.props.query}`)

    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { query } = this.props

    return (
      <>
        <NextSeo
          noindex
          title={`${query} - Phenopod`}
          description={`${query} - Phenopod`}
          canonical={`https://phenopod.com/results?query=${query}`}
          openGraph={{
            url: `https://phenopod.com/results?query=${query}`,
            type: 'website',
            title: `${query} - Phenopod`,
            description: `${query} - Phenopod`,
          }}
        />
        <div>
          <div className="-mt-1 mb-5 text-gray-700 text-lg lg:text-xl">{`Podcasts matching "${query}"`}</div>
          <ListSearchResults query={this.props.query} />
        </div>
      </>
    )
  }
}
