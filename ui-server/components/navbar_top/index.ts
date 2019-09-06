import NavbarTop, { StateToProps, DispatchToProps } from './navbar_top'
import { AppState } from '../../store'
import { getSearchBarText } from '../../selectors/ui/search'
import { Dispatch } from 'redux'
import { AppActions, SEARCH_BAR_TEXT_CHANGE } from '../../types/actions'
import { connect } from 'react-redux'

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

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(NavbarTop)
