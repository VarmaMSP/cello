import { Entity } from './app'

// Phantom types to better type selectors And keep my sanity

export type $Id<_E extends Entity> = string

export type MapById<E extends Entity> = { [id: string]: E }

export type MapOneToOne<_E1 extends Entity, E2 extends Entity> = {
  [id: string]: E2
}

export type MapOneToMany<_E1 extends Entity, _E2 extends Entity> = {
  [id: string]: string[]
}
