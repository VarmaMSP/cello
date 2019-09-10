import { connect } from 'react-redux'
import { getAllCurationIds } from 'selectors/entities/curations'
import { AppState } from 'store'
import Discover, { StateToProps } from './discover'

function mapStateToProps(state: AppState): StateToProps {
  return {
    curationIds: getAllCurationIds(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  Discover,
)
