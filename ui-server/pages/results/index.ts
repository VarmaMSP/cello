import { connect } from 'react-redux'
import { AppState } from 'store'
import ResultsPage, { OwnProps, StateToProps } from './results'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.search.searchPodcasts,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  ResultsPage,
)
