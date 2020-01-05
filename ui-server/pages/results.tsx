import { getResultsPageData } from 'actions/results'
import PageLayout from 'components/page_layout'
import SearchResultsFilter from 'components/search_results_filter/search_results_filter'
import SearchResultsList from 'components/search_results_list'
import { NextSeo } from 'next-seo'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface OwnProps {
  query: string
  sortBy: 'relevance' | 'publish_date'
  resultType: 'episode' | 'podcast'
  scrollY: number
}

export default class ResultsPage extends Component<OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store } = ctx
    const query = ctx.query['query'] as string
    const sortBy = ctx.query['sortBy'] as 'relevance' | 'publish_date'
    const resultType = ctx.query['resultType'] as 'podcast' | 'episode'

    await bindActionCreators(getResultsPageData, store.dispatch)(
      query,
      resultType,
      sortBy,
    )
  }

  componentDidMount() {
    gtag.search(this.props.query)
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { query, resultType, sortBy } = this.props

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
          <div>
            <SearchResultsFilter
              searchQuery={query}
              resultType={resultType}
              sortBy={sortBy}
            />
            <SearchResultsList
              searchQuery={query}
              resultType={resultType}
              sortBy={sortBy}
            />
          </div>
          <div />
        </PageLayout>
      </>
    )
  }
}
