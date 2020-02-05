import { getChartPageData } from 'actions/chart'
import { getHomePageData } from 'actions/home'
import ChartView from 'components/chart_view/chart_view'
import { ChartPageSeo } from 'components/seo'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

interface OwnProps {
  chartId: string
  scrollY: number
}

export default class ChartPage extends Component<OwnProps> {
  static async getInitialProps({ query, store }: PageContext): Promise<void> {
    await bindActionCreators(getHomePageData, store.dispatch)()
    await bindActionCreators(
      getChartPageData,
      store.dispatch,
    )(query['chartId'] as string)
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    return (
      <>
        <ChartPageSeo />
        <ChartView chartId={this.props.chartId} />
      </>
    )
  }
}
