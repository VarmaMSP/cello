import { connect } from 'react-redux'
import { makeGetCurrentUserHistory } from 'selectors/entities/episodes'
import { AppState } from 'store'
import ListHistory, { StateToProps } from './list_history'

function makeMapStateToProps() {
  const getCurrentUserHistory = makeGetCurrentUserHistory()

  return (state: AppState): StateToProps => ({
    history: getCurrentUserHistory(state),
  })
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  ListHistory,
)
