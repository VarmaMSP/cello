import { connect } from 'react-redux'
import { Dispatch } from 'redux'
import { getSearchBarText } from 'selectors/ui/search'
import { AppState } from 'store'
import { AppActions, SEARCH_BAR_TEXT_CHANGE } from 'types/actions'
import NavbarTop, { DispatchToProps, StateToProps } from './navbar_top'

function mapStateToProps(state: AppState): StateToProps {
  return {
    searchText: getSearchBarText(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    searchTextChange: (text: string) =>
      dispatch({ type: SEARCH_BAR_TEXT_CHANGE, text }),
  }
}

export default connect<StateToProps, DispatchToProps, {}, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(NavbarTop)
