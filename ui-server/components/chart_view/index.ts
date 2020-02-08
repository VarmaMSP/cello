import { connect } from 'react-redux'
import { getPodcastsByIds } from 'selectors/entities/podcasts'
import { AppState } from 'store'
import ChartView, { OwnProps, StateToProps } from './chart_view'

function mapStateToProps(
  state: AppState,
  { category }: OwnProps,
): StateToProps {
  return {
    podcasts: getPodcastsByIds(state, category.podcastIds),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  ChartView,
)
