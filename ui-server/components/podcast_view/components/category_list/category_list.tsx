import React from 'react'
import { Category } from 'types/models'

export interface StateToProps {
  categories: Category[]
}

export interface OwnProps {
  categoryIds: string[]
}

const CategoryList: React.FC<StateToProps & OwnProps> = ({ categories }) => {
  const byId = categories.reduce<{ [categoryId: string]: Category }>(
    (acc, c) => ({ ...acc, [c.id]: c }),
    {},
  )
  const byParentId = categories.reduce<{ [categoryId: string]: string[] }>(
    (acc, p) =>
      !!p.parentId
        ? acc
        : {
            ...acc,
            [p.id]: categories.reduce<string[]>(
              (acc, c) => (c.parentId !== p.id ? acc : [...acc, c.id]),
              [],
            ),
          },
    {},
  )

  return (
    <div className="flex flex-wrap">
      {Object.keys(byParentId).map((parentId) => {
        const parent = byId[parentId]
        const childIds = byParentId[parentId]

        return childIds.length > 0 ? (
          childIds.map((childId) => (
            <div
              id={`${parentId}${childId}`}
              className="bg-gray-300 mr-4 my-2 px-3 text-xs text-gray-800 tracking-wide leading-loose rounded-full"
            >
              <span className="font-medium">{`${parent.name}`}</span>
              <span className="mx-1">&rsaquo;</span>
              <span>{`${byId[childId].name}`}</span>
            </div>
          ))
        ) : (
          <div
            id={`${parentId}`}
            className="bg-gray-300 mr-4 my-2 px-3 text-xs font-medium text-gray-800 tracking-wide leading-loose rounded-full"
          >{`${parent.name}`}</div>
        )
      })}
    </div>
  )
}

export default CategoryList
