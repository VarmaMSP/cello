import { searchPodcasts } from 'actions/podcast'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import ResultsPage, { DispatchToProps, OwnProps, StateToProps } from './results'

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
