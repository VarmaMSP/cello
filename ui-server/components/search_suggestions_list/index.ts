import { connect } from 'react-redux'
import { getSuggestions } from 'selectors/ui/search_bar'
import { AppState } from 'store'
import SearchSuggestionsList, { StateToProps } from './search_suggestions_list'

function mapStateToProps(state: AppState): StateToProps {
  return {
    suggestions: getSuggestions(state),
  }
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  SearchSuggestionsList,
)
