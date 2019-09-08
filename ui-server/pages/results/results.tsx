import LoadingPage from 'components/loading_page'
import SearchResults from 'components/search_results'
import React, { Component } from 'react'
import { RequestState } from 'reducers/requests/utils'
import { PageContext } from 'types/utilities'

export interface StateToProps {
  reqState: RequestState
}

export interface DispatchToProps {
  loadSearchResults: (searchQuery: string) => void
}

export interface OwnProps {
  searchQuery: string
  preventIntialLoad: boolean
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

export default class ResultsPage extends Component<Props> {
  static async getInitialProps({ query }: PageContext) {
    const searchQuery = query['search_query'] as string
    return { searchQuery, preventInitalLoad: false }
  }

  componentDidUpdate(prevProps: Props) {
    if (prevProps.searchQuery != this.props.searchQuery) {
      this.props.loadSearchResults(this.props.searchQuery)
    }
  }

  componentDidMount() {
    if (!this.props.preventIntialLoad) {
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
