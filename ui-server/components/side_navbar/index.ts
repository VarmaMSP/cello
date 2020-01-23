import { connect } from 'react-redux'
import { getIsUserSignedIn } from 'selectors/session'
import { getViewportSize } from 'selectors/window'
import { AppState } from 'store'
import SideNavbar, { StateToProps } from './side_navbar'

function mapStateToProps(state: AppState): StateToProps {
  return {
    userSignedIn: getIsUserSignedIn(state),
    viewportSize: getViewportSize(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  SideNavbar,
)
