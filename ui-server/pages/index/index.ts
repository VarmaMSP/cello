import { connect } from 'react-redux'
import { AppState } from 'store'
import CurationsPage, { StateToProps } from './curations'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.curation.getAllCurations,
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  CurationsPage,
)
