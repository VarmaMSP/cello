import { connect } from 'react-redux'
import { getIsUserSignedIn } from 'selectors/session'
import { AppState } from 'store'
import HomeView, { StateToProps } from './home_view'

function mapStateToProps(state: AppState): StateToProps {
  return {
    isUserSignedIn: getIsUserSignedIn(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  HomeView,
)
