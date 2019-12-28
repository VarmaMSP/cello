import { connect } from 'react-redux'
import { makeGetSubCategories } from 'selectors/entities/curations'
import { AppState } from 'store'
import RelatedCharts, { OwnProps, StateToProps } from './related_charts'

function makeMapStateToProps() {
  const getChartChildren = makeGetSubCategories()

  return (state: AppState, { chartId }: OwnProps): StateToProps => {
    return {
      relatedCharts: getChartChildren(state, chartId),
    }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(RelatedCharts)
