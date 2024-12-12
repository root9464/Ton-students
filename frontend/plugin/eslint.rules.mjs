const customRulesPlugin = {
  rules: {
    'no-data-or-empty': {
      meta: {
        type: 'problem',
        docs: {
          description: 'Disallow usage of `data || ""` outside JSX in a return statement.',
          category: 'Best Practices',
          recommended: false,
        },
        schema: [],
      },
      create(context) {
        let insideReturnStatement = false;

        return {
          ReturnStatement() {
            insideReturnStatement = true;
          },
          'ReturnStatement:exit'() {
            insideReturnStatement = false;
          },
          LogicalExpression(node) {
            if (
              node.operator === '||' &&
              (node.left.type === 'Literal' || node.right.type === 'Literal') &&
              (node.left.value === '' || node.right.value === '')
            ) {
              if (!insideReturnStatement) {
                context.report({
                  node,
                  message: 'Avoid using `value || ""` outside of JSX in return statements.',
                });
              }
            }
          },
        };
      },
    },
  },
};

export default customRulesPlugin;
