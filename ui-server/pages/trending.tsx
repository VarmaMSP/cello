import { getTrendingPodcasts } from 'actions/podcast'
import ListTrendingPodcasts from 'components/list_trending_podcasts'
import LoadingPage from 'components/loading_page'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { AppState } from 'store'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface StateToProps {
  isLoading: boolean // FIXME
}

interface OwnProps {
  scrollY: number
}

class TrendingPage extends React.Component<StateToProps & OwnProps> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store } = ctx
    await bindActionCreators(getTrendingPodcasts, store.dispatch)()
  }

  componentDidMount() {
    gtag.pageview('/trending')

    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isLoading } = this.props

    if (isLoading) {
      return <LoadingPage />
    }

    return <ListTrendingPodcasts />
  }
}

function mapStateToProps(): StateToProps {
  return {
    isLoading: false,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  TrendingPage,
)
