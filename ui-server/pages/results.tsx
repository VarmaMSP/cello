import { getResultsPageData } from 'actions/results'
import PageLayout from 'components/page_layout'
import SearchResultsFilter from 'components/search_results_filter'
import SearchResultsList from 'components/search_results_list'
import { ResultsPageSeo } from 'components/seo'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import * as T from 'types/actions'
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
  static loadPropsIntoStore(ctx: PageContext) {
    const { store, query } = ctx
    const q = query['query'] as string
    const sortBy = query['sortBy'] as SearchSortBy
    const resultType = query['resultType'] as SearchResultType

    store.dispatch({
      type: T.SEARCH_RESULTS_QUERY,
      query: q,
    })

    store.dispatch({
      type: T.SEARCH_RESULTS_RESULT_TYPE,
      resultType,
    })

    store.dispatch({
      type: T.SEARCH_RESULTS_SORT_BY,
      sortBy,
    })

    store.dispatch({
      type: T.SEARCH_SUGGESTIONS_RESET,
    })
  }

  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store, query } = ctx

    await bindActionCreators(getResultsPageData, store.dispatch)(
      query['query'] as string,
      query['resultType'] as SearchResultType,
      query['sortBy'] as SearchSortBy,
    )
  }

  componentDidUpdate(prevProps: OwnProps) {
    const { scrollY } = this.props
    if (scrollY !== prevProps.scrollY) {
      window.window.scrollTo(0, this.props.scrollY)
    }
  }

  componentDidMount() {
    gtag.search(this.props.query)
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    return (
      <>
        <ResultsPageSeo query={this.props.query} />
        <PageLayout>
          <div className="pt-4">
            <SearchResultsFilter />
            <SearchResultsList />
          </div>
          <div />
        </PageLayout>
      </>
    )
  }
}
