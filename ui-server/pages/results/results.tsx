import { searchPodcasts } from 'actions/podcast'
import LoadingPage from 'components/loading_page'
import SearchResults from 'components/search_results'
import React, { Component } from 'react'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

export interface StateToProps {
  reqState: RequestState
}

export interface OwnProps {
  searchQuery: string
}

interface Props extends StateToProps, OwnProps {}

export default class ResultsPage extends Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<OwnProps> {
    const { query, store, isServer } = ctx
    const searchQuery = query['search_query'] as string
    const loadResults = bindActionCreators(searchPodcasts, store.dispatch)(
      searchQuery,
    )

    if (isServer) {
      await loadResults
    }
    return { searchQuery }
  }

  render() {
    const { reqState, searchQuery } = this.props

    if (reqState.status === 'STARTED') {
      return <LoadingPage />
    }

    if (reqState.status === 'SUCCESS') {
      return (
        <div>
          <div className="-mt-1 mb-5 text-gray-700 text-lg lg:text-xl">{`Podcasts matching "${searchQuery}"`}</div>
          <SearchResults searchQuery={this.props.searchQuery} />
        </div>
      )
    }

    return <></>
  }
}
