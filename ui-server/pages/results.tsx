import { getResultsPageData } from 'actions/results'
import PageLayout from 'components/page_layout'
import SearchResultsFilter from 'components/search_results_filter'
import SearchResultsList from 'components/search_results_list'
import { NextSeo } from 'next-seo'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import {
  SEARCH_BAR_UPDATE_TEXT,
  SEARCH_RESULTS_QUERY,
  SEARCH_RESULTS_RESULT_TYPE,
  SEARCH_RESULTS_SORT_BY,
} from 'types/actions'
import { SearchResultType, SearchSortBy } from 'types/search'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface OwnProps {
  query: string
  sortBy: SearchSortBy
  resultType: SearchResultType
  scrollY: number
}

export default class ResultsPage extends Component<OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store, query } = ctx

    const q = query['query'] as string
    const sortBy = query['sortBy'] as SearchSortBy
    const resultType = query['resultType'] as SearchResultType

    store.dispatch({
      type: SEARCH_BAR_UPDATE_TEXT,
      text: q,
    })

    store.dispatch({
      type: SEARCH_RESULTS_QUERY,
      query: q,
    })

    store.dispatch({
      type: SEARCH_RESULTS_RESULT_TYPE,
      resultType,
    })

    store.dispatch({
      type: SEARCH_RESULTS_SORT_BY,
      sortBy,
    })

    await bindActionCreators(getResultsPageData, store.dispatch)(
      query['query'] as string,
      query['resultType'] as SearchResultType,
      query['sortBy'] as SearchSortBy,
    )
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
          <div>
            <SearchResultsFilter />
            <SearchResultsList />
          </div>
          <div />
        </PageLayout>
      </>
    )
  }
}
