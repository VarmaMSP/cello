import { searchPodcasts } from 'actions/podcast'
import ListSearchResults from 'components/list_search_results'
import LoadingPage from 'components/loading_page'
import React, { Component } from 'react'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

export interface StateToProps {
  reqState: RequestState
}

export interface OwnProps {
  query: string
  scrollY: number
}

interface Props extends StateToProps, OwnProps {}

export default class ResultsPage extends Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { query, store, isServer } = ctx
    const loadResults = bindActionCreators(searchPodcasts, store.dispatch)(
      query['query'] as string,
    )

    if (isServer) {
      await loadResults
    }
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState, query } = this.props

    if (reqState.status === 'STARTED') {
      return <LoadingPage />
    }

    if (reqState.status === 'SUCCESS') {
      return (
        <div>
          <div className="-mt-1 mb-5 text-gray-700 text-lg lg:text-xl">{`Podcasts matching "${query}"`}</div>
          <ListSearchResults query={this.props.query} />
        </div>
      )
    }

    return <></>
  }
}
