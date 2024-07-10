# [general concepts](https://refine.dev/docs/guides-concepts/general-concepts/#ui-integrations-2)

## headless architecture

not limited by pre-styled components

refine is built up of `hooks`, `components` & `providers`
- ui/logic are decoupled

### resources

represent entities which compose the project - typically data. typical resource def:

```
export const App = () => {
  return (
    <Refine
      resources={[
        {
          name: "products",
          list: "/my-products",
          show: "/my-products/:id",
          edit: "/my-products/:id/edit",
          create: "/my-products/new",
        },
      ]}
    >
      {/* ... */}
    </Refine>
  );
};
```

### providers

used to manage aspects of the applications - e.g. routing, data fetching. they are pluggable; there are built-in providers, or you can create custom.

### hooks

hook-based architecture, refine's are headless. they are library agnostic, unified interface. different built-in providers for routers of all frameworks.

`useGo` hook exported from `@refinedev/core` 'can be used to navigate to a specific resource's page in your application regardless of your routing solution'.

# providers

## [data](https://refine.dev/docs/data/data-provider/)

bridge between data & frontend. each op is associated with a specific resource.

```
const myDataProvider: DataProvider = {
  getOne: async ({ resource, id }) => {
    const response = await fetch(
      `https://example.com/api/v1/${resource}/${id}`,
    );
    const data = await response.json();

    return { data };
  },
  // other methods...
};
```

## hooks

`useList`, `useOne`, `useCreate`, `useEdit`, `useShow` are all used to fetch various component data.

# auth providers

```
export const authProvider: AuthProvider = {
  login: async ({ email, password }) => {
    const { status } = handleLogin(email, password);

    if (status === 200) {
      return { success: true, redirectTo: "/dashboard" };
    } else {
      return {
        success: false,
        error: { name: "Login Error", message: "Invalid credentials" },
      };
    }
  },
  check: async (params) => ({}),
  logout: async (params) => ({}),
  onError: async (params) => ({}),
  register: async (params) => ({}),
  forgotPassword: async (params) => ({}),
  updatePassword: async (params) => ({}),
  getPermissions: async (params) => ({}),
  getIdentity: async (params) => ({}),
};
```

## [components](https://refine.dev/docs/guides-concepts/authentication/#components)

the `Authenticated` component from `@refinedev/core` to protect routes, components & auth.

```

const MyPage = () => (
  <Authenticated>
    // Only authenticated users can see this.
    <MyComponent />
  </Authenticated>
);
```

## hooks

`useGetIdentity` gets current user

```
export const DashboardPage = () => {
  const {
    data: { name },
  } = useGetIdentity();

  return <>Welcome {name}!</>;
};
```

## ui integrations

pre-built components render auth provider out-of-the-box. their layouts can render user info, logout buttons, etc.

other components include `AuthPage` for integrations `Login`, `Register`, `Forgot Password`, `Reset Password` & more.

## access control provider

manages what users can access within the app based on their perms, uses resource definition to determine access rights.

```
const myAccessControlProvider: AccessControlProvider = {
  can: async ({ resource, action }) => {
    if (resource === "users" && action === "block") {
      return { can: false };
    }

    return { can: true };
  },
};

export const App = () => {
  return (
    <Refine accessControlProvider={myAccessControlProvider}>{/* ... */}</Refine>
  );
};
```

## components

wrap `CanAccess` component to wrap relevant parts of the app to control access.

```
export const MyPage = () => {
  return (
    <CanAccess resource="users" action="show" params={{ id: 1 }}>
      <>
        My Page
        <CanAccess
          resource="users"
          action="block"
          params={{ id: 1 }}
          fallback={"You are not authorized."}
        >
          // Only authorized users can see this button.
          <BlockUserButton />
        </CanAccess>
      </>
    </CanAccess>
  );
};
```

## hooks

use `useCan` to control access in your components.

```
export const MyPage = () => {
  const { data: show } = useCan({
    resource: "users",
    action: "show",
    params: { id: 1 },
  });
  const { data: block } = useCan({
    resource: "users",
    action: "block",
    params: { id: 1 },
  });

  if (!show?.can) {
    return <ErrorComponent />;
  }

  return (
    <>
      My Page
      {block?.can && <BlockUserButton />}
      {!block?.can && "You are not authorized."}
    </>
  );
};
```

## ui integrations

```
export const MyPage = () => {
  return (
    <>
      My Page
      {/* Only authorized users can see this button. */}
      <DeleteButton resource="users" recordItemId={1} />
    </>
  );
};
```

This applies to all buttons like CreateButton, EditButton, ShowButton, ListButton.

# notification provider

refine can automatically show notifications for CRUD operations & errors.

## hooks

`data`, `mutation` & `auth` hooks can automatically show notifications for actions & errors + modify these notifications per hook.

```
export const MyPage = () => {
  const { mutate } = useDelete();

  return (
    <Button
      onClick={() => {
        mutate({
          resource: "products",
          id: 1,
          successNotification: () => ({
            message: "Product Deleted",
            description: "Product has been deleted successfully.",
            type: "success",
          }),
          errorNotification: () => ({
            message: "Product Delete Error",
            description: "An error occurred while deleting the product.",
            type: "error",
          }),
        });
      }}
    >
      Delete Product
    </Button>
  );
};
```

`useNotification` shows notifications in your application.

```
export const MyPage = () => {
  const { open, close } = useNotification();

  return (
    <>
      <Button
        onClick={() => {
          open?.({
            key: "my-notification",
            message: "Test Notification",
            description: "This is a test notification.",
            type: "success", // success | error | progress
          });
        }}
      >
        Show notification
      </Button>
      <Button
        onClick={() => {
          close?.("my-notification");
        }}
      >
        Close Notification
      </Button>
    </>
  );
};
```

# [i18n](https://refine.dev/docs/guides-concepts/i18n/)

# router provider

router provider helps the app understand the relationship between resources & routes. enables nav features like breadcrumbs, automatic redirections, rendering menu items, inferring hook params, etc.


