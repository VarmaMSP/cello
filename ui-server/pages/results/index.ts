import { AppState } from '../../store'
import ResultsPage, { StateToProps, DispatchToProps, OwnProps } from './results'
import { Dispatch, bindActionCreators } from 'redux'
import { AppActions } from '../../types/actions'
import { searchPodcasts } from '../../actions/podcast'
import { connect } from 'react-redux'

function mapStateToProps(state: AppState): StateToProps {
  return {
    reqState: state.requests.search.searchPodcasts,
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadSearchResults: bindActionCreators(searchPodcasts, dispatch),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(ResultsPage)
