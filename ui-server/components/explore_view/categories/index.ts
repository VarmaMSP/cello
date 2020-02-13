import { connect } from 'react-redux'
import { makeGetPrimaryCategories } from 'selectors/entities/categories'
import { AppState } from 'store'
import Categories, { StateToProps } from './categories'

function mapStateToProps() {
  const getPrimaryCategories = makeGetPrimaryCategories()

  return (state: AppState): StateToProps => ({
    categories: getPrimaryCategories(state),
  })
}

export default connect<StateToProps, {}, {}, AppState>(mapStateToProps)(
  Categories,
)
