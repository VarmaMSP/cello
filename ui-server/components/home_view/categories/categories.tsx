import CategoryList from 'components/category_list'
import React from 'react'
import { Curation } from 'types/models'

export interface StateToProps {
  categories: Curation[]
}

const Categories: React.FC<StateToProps> = () => {
  return (
    <>
      <h1 className="pb-1 text-xl tracking-wide text-gray-800 font-medium">
        {'Browse By Categories'}
      </h1>
      <hr className="mb-3" />
      <div className="flex flex-col flex-wrap" style={{ height: '50rem' }}>
        <CategoryList className="w-1/4" />
      </div>
    </>
  )
}

export default Categories
