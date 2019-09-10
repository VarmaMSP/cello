import { getCurations } from 'actions/curations'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import CurationsPage, {
  DispatchToProps,
  OwnProps,
  StateToProps,
} from './curations'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.curation.getAllCurations,
  }
}

function mapDispatchToProps(dispatch: Dispatch): DispatchToProps {
  return {
    loadCurations: bindActionCreators(getCurations, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(CurationsPage)
