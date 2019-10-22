import { searchPodcasts } from 'actions/podcast'
import ListSearchResults from 'components/list_search_results'
import LoadingPage from 'components/loading_page'
import { NextSeo } from 'next-seo'
import React, { Component } from 'react'
import { connect } from 'react-redux'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { AppState } from 'store'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface StateToProps {
  reqState: RequestState
}

interface OwnProps {
  query: string
  scrollY: number
}

class ResultsPage extends Component<StateToProps & OwnProps> {
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
    gtag.pageview(`/results?query=${this.props.query}`)

    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState, query } = this.props

    if (reqState.status === 'STARTED') {
      return <LoadingPage />
    }
    if (reqState.status === 'SUCCESS') {
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
    return <></>
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.search.searchPodcasts,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  ResultsPage,
)
