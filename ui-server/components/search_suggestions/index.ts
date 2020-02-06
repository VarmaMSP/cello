import { connect } from 'react-redux'
import { AppState } from 'store'
import SearchSuggestions, { StateToProps } from './search_suggestions'

function mapStateToProps(state: AppState): StateToProps {
  return {
    podcasts: state.entities.searchSuggestions.podcasts,
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  SearchSuggestions,
)
