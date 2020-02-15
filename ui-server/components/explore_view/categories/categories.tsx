import { ChartLink } from 'components/link'
import { Category } from 'models'
import React from 'react'

export interface StateToProps {
  categories: Category[]
}

const Categories: React.FC<StateToProps> = ({ categories }) => {
  return (
    <>
      <h1 className="pb-1 text-xl tracking-wide font-semibold">
        {'Browse By Categories'}
      </h1>
      <hr className="mb-3" />
      <div className="flex flex-wrap">
        {categories.map((x) => (
          <div
            key={x.id}
            className="w-full md:w-1/4 px-2 py-2 text-gray-900 hover:text-black font-medium tracking-wide"
          >
            <ChartLink chartUrlParam={x.urlParam}>
              <a>{x.name}</a>
            </ChartLink>
          </div>
        ))}
      </div>
    </>
  )
}

export default Categories
