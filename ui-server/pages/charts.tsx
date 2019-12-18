import { getChartPageData } from 'actions/chart'
import { getHomePageData } from 'actions/home'
import ChartView from 'components/chart_view/chart_view'
import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { PageContext } from 'types/utilities'

interface OwnProps {
  chartId: string
  scrollY: number
}

export default class ChartPage extends Component<OwnProps> {
  static async getInitialProps({ query, store }: PageContext): Promise<void> {
    await bindActionCreators(
      getChartPageData,
      store.dispatch,
    )(query['chartId'] as string)

    await bindActionCreators(getHomePageData, store.dispatch)()
  }

  componentDidMount() {
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
