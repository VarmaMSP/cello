import { connect } from 'react-redux'
import { AppState } from 'store'
import Categories, { StateToProps } from './categories'

function makeMapStateToProps() {
  return (): StateToProps => {
    return { categories: [] }
  }
}

export default connect<StateToProps, {}, {}, AppState>(makeMapStateToProps())(
  Categories,
)
