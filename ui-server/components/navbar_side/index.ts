import { connect } from 'react-redux'
import { getCurrentUrlPath } from 'selectors/browser/urlPath'
import { getIsUserSignedIn } from 'selectors/entities/users'
import { AppState } from 'store'
import Navbar, { StateToProps } from './navbar_side'

function mapStateToProps(state: AppState): StateToProps {
  return {
    userSignedIn: getIsUserSignedIn(state),
    currentUrlPath: getCurrentUrlPath(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(Navbar)
