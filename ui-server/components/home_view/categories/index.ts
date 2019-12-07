import { connect } from 'react-redux'
import { makeGetCategories } from 'selectors/entities/podcast_lists'
import { AppState } from 'store'
import Categories, { StateToProps } from './categories'

function makeMapStateToProps() {
  const getCategories = makeGetCategories()

  return (state: AppState): StateToProps => {
    return {
      categories: getCategories(state),
    }
  }
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Categories,
)
