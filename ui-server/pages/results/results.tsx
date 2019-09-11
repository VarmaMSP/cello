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

export interface DispatchToProps {
  loadSearchResults: (searchQuery: string) => void
}

export interface OwnProps {
  searchQuery: string
  preventInitialLoad: boolean
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

export default class ResultsPage extends Component<Props> {
  static async getInitialProps({
    query,
    isServer,
    store,
  }: PageContext): Promise<OwnProps> {
    const loadResults = bindActionCreators(searchPodcasts, store.dispatch)
    const searchQuery = query['search_query'] as string

    await loadResults(searchQuery)
    return { searchQuery, preventInitialLoad: isServer }
  }

  componentDidUpdate(prevProps: Props) {
    if (prevProps.searchQuery != this.props.searchQuery) {
      this.props.loadSearchResults(this.props.searchQuery)
    }
  }

  componentDidMount() {
    if (!this.props.preventInitialLoad) {
      this.props.loadSearchResults(this.props.searchQuery)
    }
  }

  render() {
    const { reqState, searchQuery } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }

    return (
      <div>
        <div className="-mt-1 mb-5 text-gray-700 text-lg lg:text-xl">{`Podcasts matching "${searchQuery}"`}</div>
        <SearchResults searchQuery={this.props.searchQuery} />
      </div>
    )
  }
}
