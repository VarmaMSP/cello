import { connect } from 'react-redux'
import { AppState } from 'store'
import CurationsPage, { OwnProps, StateToProps } from './curations'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.curation.getAllCurations,
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  CurationsPage,
)
