import { connect } from 'react-redux'
import { getAllCategories } from 'selectors/entities/categories'
import { AppState } from 'store'
import CategoryList, { OwnProps, StateToProps } from './category_list'

function mapStateToProps(state: AppState): StateToProps {
  return {
    categories: getAllCategories(state),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  CategoryList,
)
