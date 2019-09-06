import React, { Component } from 'react'
import { PageContext } from '../../types/utilities'
import { RequestState } from '../../reducers/requests/utils'
import SearchResults from '../../components/search_results'

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
    const { reqState } = this.props

    if (reqState.status == 'STARTED') {
      return <>Loading</>
    }
    return <SearchResults searchQuery={this.props.searchQuery} />
  }
}
