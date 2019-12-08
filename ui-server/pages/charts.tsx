import { getPodcastsInChart } from 'actions/chart'
import { getHomePageData } from 'actions/home'
import ChartView from 'components/chart_view/chart_view'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'
import * as gtag from 'utils/gtag'

interface OwnProps {
  chartId: string
  scrollY: number
}

export default class ChartsPage extends Component<OwnProps> {
  static async getInitialProps({ query, store }: PageContext): Promise<void> {
    await bindActionCreators(getHomePageData, store.dispatch)()
    await bindActionCreators(
      getPodcastsInChart,
      store.dispatch,
    )(query['chartId'] as string)
  }

  componentDidMount() {
    gtag.pageview(`/charts/${this.props.chartId}`)
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    return (
      <div>
        <ChartView chartId={this.props.chartId} />
      </div>
    )
  }
}
