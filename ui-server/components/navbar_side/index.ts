import { connect } from 'react-redux'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { AppState } from 'store'
import Navbar, { StateToProps } from './navbar_side'

function mapStateToProps(state: AppState): StateToProps {
  return {
    userSignedIn: getIsUserSignedIn(state),
    currentPathName: state.ui.currentPathName,
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(Navbar)
