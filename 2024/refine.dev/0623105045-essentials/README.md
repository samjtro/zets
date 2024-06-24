# essentials

## data fetching

`<Refine />` component accepts `dataProvider` prop which is used to handle all the data fetching and mutation operations.

`src/providers/data-provider.ts` contains all the methods needed to implement the data provider.

sample empty data provider:

```
import type { DataProvider } from "@refinedev/core";

const API_URL = "https://api.fake-rest.refine.dev";

export const dataProvider: DataProvider = {
  getOne: () => {
    throw new Error("Not implemented");
  },
  update: () => {
    throw new Error("Not implemented");
  },
  getList: () => {
    throw new Error("Not implemented");
  },
  create: () => {
    throw new Error("Not implemented");
  },
  deleteOne: () => {
    throw new Error("Not implemented");
  },
  getApiUrl: () => API_URL,
  // Optional methods:
  // getMany: () => { /* ... */ },
  // createMany: () => { /* ... */ },
  // deleteMany: () => { /* ... */ },
  // updateMany: () => { /* ... */ },
  // custom: () => { /* ... */ },
};
```

then, add the following to `src/App.tsx`:

```
import { dataProvider } from "./providers/data-provider";
... App() {
    return (
        <Refine dataProvider={dataProvider}>
...
```

## fetching a record
