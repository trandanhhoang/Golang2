# from functools import reduce
# class ASTGeneration(MPVisitor):

#     def visitProgram(self, ctx: MPParser.ProgramContext):
#         varDeclList = sum([self.visit(x) for x in ctx.vardecl()])
#         return Program(varDeclList)

#     def visitVardecl(self, ctx: MPParser.ExpContext):
#         idList = self.visit(ctx.idList())
#         idType = self.visit(ctx.typ())
#         listVarDecl =  [VarDecl(id[1],id[0],idType,None) for id in idList]
#         return listVarDecl

#     def visitIdlsit(self, ctx: MPParser.TermContext):
#         listMem = [self.visit(x) for x in ctx.mem()]
#         return listMem

#     def visitMem(self, ctx: MPParser.FactorContext):
#         static = false
#         if ctx.STATIC():
#             static = true
#         id = Id(ctx.ID().getText())
#         return [static,id]

#     def visitTyp(self, ctx: MPParser.OperandContext):
#         return IntType() if ctx.INT() else FloatType()

a= [1,2,3]
b= [4,5,6]
c= [a,b]
print(sum([x for x in c],[]))
