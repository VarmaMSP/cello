import { getResultsPageData } from 'actions/results'
import PageLayout from 'components/page_layout'
import SearchResultsList from 'components/search_results_list'
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
      getResultsPageData,
      store.dispatch,
    )(query['query'] as string)
  }

  componentDidMount() {
    gtag.search(this.props.query)
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
        <PageLayout>
          <SearchResultsList searchQuery={query} />
          <div />
        </PageLayout>
      </>
    )
  }
}
