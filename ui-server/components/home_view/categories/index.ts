import { connect } from 'react-redux'
import { getMainCategories } from 'selectors/entities/curations'
import { AppState } from 'store'
import Categories, { StateToProps } from './categories'

function makeMapStateToProps() {
  return (state: AppState): StateToProps => {
    return { categories: getMainCategories(state) }
  }
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Categories,
)
