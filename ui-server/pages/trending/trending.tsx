import { getTrendingPodcasts } from 'actions/podcast'
import ListTrendingPodcasts from 'components/list_trending_podcasts'
import LoadingPage from 'components/loading_page'
import React from 'react'
import { RequestState } from 'reducers/requests/utils'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

export interface StateToProps {
  reqState: RequestState
}

export interface OwnProps {
  scrollY: number
}

interface Props extends StateToProps, OwnProps {}

export default class extends React.Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store, isServer } = ctx
    const loadTrendingPodcasts = bindActionCreators(
      getTrendingPodcasts,
      store.dispatch,
    )()

    if (isServer) {
      await loadTrendingPodcasts
    }
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { reqState } = this.props

    if (reqState.status == 'STARTED') {
      return <LoadingPage />
    }

    if (reqState.status == 'SUCCESS') {
      return <ListTrendingPodcasts />
    }

    return <></>
  }
}
