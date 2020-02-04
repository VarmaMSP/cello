import { connect } from 'react-redux'
import { getCategoriesByIds } from 'selectors/entities/categories'
import { AppState } from 'store'
import CategoryList, { OwnProps, StateToProps } from './category_list'

function mapStateToProps(
  state: AppState,
  { categoryIds }: OwnProps,
): StateToProps {
  return {
    categories: getCategoriesByIds(state, categoryIds),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(mapStateToProps)(
  CategoryList,
)
