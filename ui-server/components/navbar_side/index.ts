import { connect } from 'react-redux'
import { AppState } from 'store'
import Navbar, { StateToProps } from './navbar_side'

function mapStateToProps(state: AppState): StateToProps {
  return {
    currentPathName: state.ui.currentPathName,
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(Navbar)
