# Unit Testing Patterns & Standards

## 1. Scope & Isolation
- **Target**: A single logical unit (function, class, or small group of tightly coupled components).
- **Rule**: If your unit test requires a real database or network, it is NOT a unit test.
- **Speed**: Unit tests must run in milliseconds.

---

## 2. The AAA Pattern (Arrange-Act-Assert)
Every unit test should follow this structure:

- **Arrange**: Set up the initial state, mocks, and input data.
- **Act**: Call the function or method being tested.
- **Assert**: Verify that the output or state change matches expectation.

```typescript
// Example
it('should calculate total price with tax', () => {
    // Arrange
    const cart = new Cart();
    cart.add({ price: 100 });
    
    // Act
    const total = cart.calculateTotal(0.1); // 10% tax
    
    // Assert
    expect(total).toBe(110);
});
```

---

## 3. What to Test
- **Happy Path**: The most common and successful usage of the code.
- **Edge Cases**: Null inputs, empty strings, maximum values, boundary conditions.
- **Error Handling**: Does the code throw or return the correct error when it should?

---

## 4. Best Practices
- **One Logical Assertion per Test**: Avoid testing 10 different things in one `it` block.
- **Descriptive Naming**: Use `should [behavior] when [context]` naming style.
- **Avoid Implementation Details**: Assert on the *result*, not how the code arrived at that result.
- **Deterministic**: Tests must not rely on current time, random numbers, or global state without proper control/mocking.
