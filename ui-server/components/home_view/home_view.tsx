import React from 'react'
import Categories from './categories'
import Recommended from './recommended'

const HomeView: React.FC<{}> = () => {
  return (
    <div>
      <Recommended />
      <Categories />
    </div>
  )
}

export default HomeView
